package api

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ellcrys/util"
	"github.com/ncodes/cocoon/core/api/api/proto"
	"github.com/ncodes/cocoon/core/common"
	orderer_proto "github.com/ncodes/cocoon/core/orderer/proto"
	"github.com/ncodes/cocoon/core/types"
	"github.com/ncodes/cstructs"
	context "golang.org/x/net/context"
)

var (
	// CocoonStatusCreated indicates a created cocoon
	CocoonStatusCreated = "created"

	// CocoonStatusStarted indicates a started cocoon cocoon
	CocoonStatusStarted = "started"

	// CocoonStatusRunning indicates a running cocoon code
	CocoonStatusRunning = "running"

	// CocoonStatusStopped indicates a stopped cocoon
	CocoonStatusStopped = "stopped"
)

// makeCocoonKey constructs a cocoon key
func (api *API) makeCocoonKey(id string) string {
	return fmt.Sprintf("cocoon.%s", id)
}

// updateCocoon creates a new cocoon by overriding an existing one (if any).
func (api *API) updateCocoon(ctx context.Context, cocoon *types.Cocoon) error {
	var protoCreateCocoonReq proto.CreateCocoonRequest
	if err := cstructs.Copy(cocoon, &protoCreateCocoonReq); err != nil {
		return err
	}
	protoCreateCocoonReq.OptionAllowDuplicate = true
	_, err := api.CreateCocoon(ctx, &protoCreateCocoonReq)
	if err != nil {
		return err
	}
	return nil
}

// updateCocoonStatusOnStarted checks on interval the cocoon status and update the
// status field when the cocoon status is `started`. It returns immediately
// an error is encountered. The function will block till success or failure.
func (api *API) updateCocoonStatusOnStarted(ctx context.Context, cocoon *types.Cocoon) error {
	for {
		status, err := api.GetCocoonStatus(cocoon.ID)
		if err != nil {
			return err
		}
		if status == CocoonStatusRunning {
			cocoon.Status = CocoonStatusStarted
			return api.updateCocoon(ctx, cocoon)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// CreateCocoon creates a cocoon
func (api *API) CreateCocoon(ctx context.Context, req *proto.CreateCocoonRequest) (*proto.Response, error) {

	var err error
	var claims jwt.MapClaims

	if claims, err = api.checkCtxAccessToken(ctx); err != nil {
		return nil, types.ErrInvalidOrExpiredToken
	}

	var cocoon types.Cocoon
	cstructs.Copy(req, &cocoon)
	allowDup := req.OptionAllowDuplicate
	timeCreated, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)
	req = nil

	if err := ValidateCocoon(&cocoon); err != nil {
		return nil, err
	}

	loggedInIdentity := claims["identity"].(string)

	// if cocoon has a identity set, ensure the identity matches that of the logged in user
	if len(cocoon.IdentityID) != 0 && cocoon.IdentityID != loggedInIdentity {
		return nil, fmt.Errorf("Permission denied: You do not have permission to perform this operation")
	}

	// set identity id of cocoon to the identity of the logged in user
	if len(cocoon.IdentityID) == 0 {
		cocoon.IdentityID = loggedInIdentity
	}

	if len(cocoon.Status) == 0 {
		cocoon.Status = CocoonStatusCreated
	}

	// add cocoon owner identity if not included
	if !util.InStringSlice(cocoon.Signatories, cocoon.IdentityID) {
		cocoon.Signatories = append(cocoon.Signatories, cocoon.IdentityID)
	}

	if !allowDup {
		_, err = api.GetCocoon(ctx, &proto.GetCocoonRequest{
			ID: cocoon.ID,
		})

		if err != nil && err != types.ErrCocoonNotFound {
			return nil, err
		} else if err != types.ErrCocoonNotFound {
			return nil, fmt.Errorf("cocoon with matching ID already exists")
		}
	}

	// if a link cocoon id is provided, check if the linked cocoon exists
	if len(cocoon.Link) > 0 {
		_, err = api.GetCocoon(ctx, &proto.GetCocoonRequest{
			ID: cocoon.Link,
		})
		if err != nil && err != types.ErrCocoonNotFound {
			return nil, err
		} else if err == types.ErrCocoonNotFound {
			return nil, fmt.Errorf("cannot link to a non-existing cocoon")
		}
	}

	ordererConn, err := api.ordererDiscovery.GetGRPConn()
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	value := cocoon.ToJSON()
	odc := orderer_proto.NewOrdererClient(ordererConn)
	_, err = odc.Put(ctx, &orderer_proto.PutTransactionParams{
		CocoonID:   "",
		LedgerName: types.GetGlobalLedgerName(),
		Transactions: []*orderer_proto.Transaction{
			&orderer_proto.Transaction{
				Id:        util.UUID4(),
				Key:       api.makeCocoonKey(cocoon.ID),
				Value:     string(value),
				CreatedAt: timeCreated.Unix(),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		Status: 200,
		Body:   value,
	}, nil
}

// UpdateCocoon updates a cocoon and optionally creates a new
// release. A new release is created when Release fields are
// set/defined. No release is created if no change was made to previous release.
func (api *API) UpdateCocoon(ctx context.Context, req *proto.UpdateCocoonRequest) (*proto.Response, error) {

	var claims jwt.MapClaims

	resp, err := api.GetCocoon(ctx, &proto.GetCocoonRequest{ID: req.ID})
	if err != nil && err != types.ErrCocoonNotFound {
		return nil, err
	}

	var cocoon types.Cocoon
	util.FromJSON(resp.Body, &cocoon)

	if claims, err = api.checkCtxAccessToken(ctx); err != nil {
		return nil, types.ErrInvalidOrExpiredToken
	}

	loggedInIdentity := claims["identity"].(string)

	// if cocoon has a identity set, ensure the identity matches that of the logged in user
	if len(cocoon.IdentityID) != 0 && cocoon.IdentityID != loggedInIdentity {
		return nil, fmt.Errorf("Permission denied: You do not have permission to perform this operation")
	}

	// update new non-release fields that are different
	// from their existing value
	cocoonUpdated := false
	if len(req.Memory) > 0 && req.Memory != cocoon.Memory {
		cocoon.Memory = req.Memory
		cocoonUpdated = true
	}
	if len(req.CPUShares) > 0 && req.CPUShares != cocoon.CPUShares {
		cocoon.CPUShares = req.CPUShares
		cocoonUpdated = true
	}
	if req.NumSignatories > 0 && req.NumSignatories != cocoon.NumSignatories {
		cocoon.NumSignatories = req.NumSignatories
		cocoonUpdated = true
	}
	if req.SigThreshold > 0 && req.SigThreshold != cocoon.SigThreshold {
		cocoon.SigThreshold = req.SigThreshold
		cocoonUpdated = true
	}

	// validate updated cocoon
	if err = ValidateCocoon(&cocoon); err != nil {
		return nil, err
	}

	// get the last deployed release. if no recently deployed release,
	// get the initial release
	var recentReleaseID = cocoon.LastDeployedRelease
	if recentReleaseID == "" {
		recentReleaseID = cocoon.Releases[len(cocoon.Releases)-1]
	}

	resp, err = api.GetRelease(ctx, &proto.GetReleaseRequest{ID: recentReleaseID})
	if err != nil {
		return nil, err
	}

	var release types.Release
	util.FromJSON(resp.Body, &release)

	// Create new release and set values if any of the release field changed
	var releaseUpdated = false
	if len(req.URL) > 0 && req.URL != release.URL {
		release.URL = req.URL
		releaseUpdated = true
	}
	if len(req.ReleaseTag) > 0 && req.ReleaseTag != release.ReleaseTag {
		release.ReleaseTag = req.ReleaseTag
		releaseUpdated = true
	}
	if len(req.Language) > 0 && req.Language != release.Language {
		release.Language = req.Language
		releaseUpdated = true
	}
	if len(req.BuildParam) > 0 && req.BuildParam != release.BuildParam {
		release.BuildParam = req.BuildParam
		releaseUpdated = true
	}
	if len(req.Link) > 0 && req.Link != release.Link {
		release.Link = req.Link
		releaseUpdated = true
	}

	var finalResp = map[string]interface{}{
		"newReleaseID":  "",
		"cocoonUpdated": cocoonUpdated,
	}

	// create new release if a field was changed
	if releaseUpdated {

		// reset
		release.ID = util.UUID4()
		release.VotersID = []string{}
		release.CreatedAt = time.Now().UTC().Format(time.RFC3339Nano)

		// add id to cocoon's releases
		cocoon.Releases = append(cocoon.Releases, release.ID)

		// validate release
		if err = ValidateRelease(&release); err != nil {
			return nil, err
		}

		// persist release
		var protoCreateReleaseReq proto.CreateReleaseRequest
		cstructs.Copy(release, &protoCreateReleaseReq)
		_, err = api.CreateRelease(ctx, &protoCreateReleaseReq)
		if err != nil {
			return nil, err
		}

		finalResp["newReleaseID"] = release.ID
	}

	// update cocoon if cocoon was changed or release was updated/created
	if cocoonUpdated || releaseUpdated {
		var protoCreateCocoonReq proto.CreateCocoonRequest
		cstructs.Copy(cocoon, &protoCreateCocoonReq)
		protoCreateCocoonReq.OptionAllowDuplicate = true
		_, err = api.CreateCocoon(ctx, &protoCreateCocoonReq)
		if err != nil {
			return nil, err
		}
	}

	value, _ := util.ToJSON(&finalResp)

	return &proto.Response{
		Status: 200,
		Body:   value,
	}, nil
}

// GetCocoon fetches a cocoon
func (api *API) GetCocoon(ctx context.Context, req *proto.GetCocoonRequest) (*proto.Response, error) {

	ordererConn, err := api.ordererDiscovery.GetGRPConn()
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	odc := orderer_proto.NewOrdererClient(ordererConn)
	tx, err := odc.Get(ctx, &orderer_proto.GetParams{
		CocoonID: "",
		Key:      api.makeCocoonKey(req.GetID()),
		Ledger:   types.GetGlobalLedgerName(),
	})

	if err != nil && common.CompareErr(err, types.ErrTxNotFound) != 0 {
		return nil, err
	} else if err != nil && common.CompareErr(err, types.ErrTxNotFound) == 0 {
		return nil, types.ErrCocoonNotFound
	}

	return &proto.Response{
		Status: 200,
		Body:   []byte(tx.GetValue()),
	}, nil
}

// GetCocoonStatus fetches the cocoon status.It queries the scheduler
// to find out the current service status for the cocoon.
func (api *API) GetCocoonStatus(cocoonID string) (string, error) {

	s, err := api.scheduler.GetServiceDiscoverer().GetByID("cocoon", map[string]string{
		"tag": cocoonID,
	})
	if err != nil {
		apiLog.Errorf("failed to query cocoon service status: %s", err.Error())
		return "", fmt.Errorf("failed to query cocoon service status")
	}

	if len(s) == 0 {
		return CocoonStatusStopped, nil
	}

	return CocoonStatusRunning, nil
}

// StopCocoon stops a running cocoon and sets its status to `stopped`
func (api *API) StopCocoon(ctx context.Context, req *proto.StopCocoonRequest) (*proto.Response, error) {

	var claims jwt.MapClaims
	var err error

	if claims, err = api.checkCtxAccessToken(ctx); err != nil {
		return nil, types.ErrInvalidOrExpiredToken
	}

	ordererConn, err := api.ordererDiscovery.GetGRPConn()
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	odc := orderer_proto.NewOrdererClient(ordererConn)
	tx, err := odc.Get(ctx, &orderer_proto.GetParams{
		CocoonID: "",
		Key:      api.makeCocoonKey(req.GetID()),
		Ledger:   types.GetGlobalLedgerName(),
	})

	if err != nil {
		return nil, err
	}

	var cocoon types.Cocoon
	if err = util.FromJSON([]byte(tx.GetValue()), &cocoon); err != nil {
		return nil, common.JSONCoerceErr("cocoon", err)
	}

	// ensure session identity matches cocoon identity
	if claims["identity"].(string) != cocoon.IdentityID {
		return nil, fmt.Errorf("Permission denied: You do not have permission to perform this operation")
	}

	err = api.scheduler.Stop(req.GetID())
	if err != nil {
		apiLog.Error(err.Error())
		return nil, fmt.Errorf("failed to stop cocoon")
	}

	cocoon.Status = CocoonStatusStopped

	if err = api.updateCocoon(ctx, &cocoon); err != nil {
		apiLog.Error(err.Error())
		return nil, fmt.Errorf("failed to update cocoon status")
	}

	return &proto.Response{
		Status: 200,
		Body:   []byte("done"),
	}, nil
}
