package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"strings"

	"github.com/ellcrys/cocoon/core/api/api/proto_api"
	"github.com/ellcrys/cocoon/core/common"
	"github.com/ellcrys/cocoon/core/types"
	"github.com/ellcrys/util"
	"google.golang.org/grpc"
)

// AddVote adds a new vote to a release. If isCocoonID is true,
// the id is taken to be a cocoon id and as such the vote is added
// to the latest release. A positive vote is denoted with 1 or 0 for
// negative.
func AddVote(id string, vote int, isCocoonID bool) error {

	var releaseID = id
	var cocoon types.Cocoon

	stopSpinner := util.Spinner("Please wait")
	defer stopSpinner()
	conn, err := grpc.Dial(APIAddress, grpc.WithInsecure())
	if err != nil {
		stopSpinner()
		return fmt.Errorf("unable to connect to the platform")
	}
	defer conn.Close()

	conn, err = GetAPIConnection()
	if err != nil {
		return fmt.Errorf("unable to connect to the platform")
	}
	defer conn.Close()

	ctx, cc := context.WithTimeout(context.Background(), ContextTimeout)
	defer cc()

	cl := proto_api.NewAPIClient(conn)

	// if id is a cocoon id, get the cocoon's most recent release
	if isCocoonID {

		resp, err := cl.GetCocoon(ctx, &proto_api.GetCocoonRequest{ID: id})
		if err != nil {
			stopSpinner()
			if common.CompareErr(err, types.ErrCocoonNotFound) == 0 {
				return fmt.Errorf("%s: cocoon does not exists", common.GetShortID(id))
			}
			return err
		}

		util.FromJSON(resp.Body, &cocoon)

		// set id to latest release
		releaseID = cocoon.Releases[len(cocoon.Releases)-1]
	}

	stopSpinner()
	time.Sleep(100 * time.Millisecond)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Are you sure you? Y/n: ")
		text, _ := reader.ReadString('\n')
		v := strings.TrimSpace(strings.ToLower(text))
		if v == "n" {
			log.Info("Aborted")
			return nil
		}
		if v == "y" {
			break
		}
	}

	stopSpinner = util.Spinner("Please wait")

	_, err = cl.AddVote(ctx, &proto_api.AddVoteRequest{
		ReleaseID: releaseID,
		Vote:      int32(vote),
	})
	if err != nil {
		stopSpinner()
		if common.CompareErr(err, fmt.Errorf("release not found")) == 0 {
			return fmt.Errorf("release (%s) was not found", common.GetShortID(releaseID))
		}
		return err
	}

	stopSpinner()

	voteTxt := ""
	switch vote {
	case 1:
		voteTxt = "Approve"
	case 0:
		voteTxt = "Deny"
	}

	log.Info("==> You have successfully voted")
	log.Infof("==> You voted: %d (%s)", vote, voteTxt)

	return nil
}

// GetReleases fetches one or more releases and logs them
func GetReleases(ids []string) error {

	if len(ids) > MaxBulkObjCount {
		return fmt.Errorf("max number of objects exceeded. Expects a maximum of %d", MaxBulkObjCount)
	}

	var releases []types.Release
	var err error
	var resp *proto_api.Response

	conn, err := GetAPIConnection()
	if err != nil {
		return fmt.Errorf("unable to connect to the platform")
	}
	defer conn.Close()

	for _, id := range ids {
		stopSpinner := util.Spinner("Please wait")

		ctx, cc := context.WithTimeout(context.Background(), ContextTimeout)
		defer cc()
		cl := proto_api.NewAPIClient(conn)
		resp, err = cl.GetRelease(ctx, &proto_api.GetReleaseRequest{
			ID: id,
		})
		if err != nil {
			if common.CompareErr(err, types.ErrTxNotFound) == 0 {
				stopSpinner()
				err = fmt.Errorf("No such object: %s", id)
				break
			}
			stopSpinner()
			break
		}

		var release types.Release
		if err = util.FromJSON(resp.Body, &release); err != nil {
			return common.JSONCoerceErr("cocoon", err)
		}

		releases = append(releases, release)
		stopSpinner()
	}

	if len(releases) > 0 {
		bs, _ := json.MarshalIndent(releases, "", "   ")
		log.Infof("%s", bs)
	}
	if err != nil {
		return err
	}

	return nil
}
