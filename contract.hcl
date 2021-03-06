# A contract file describes a contract specification
contracts {
    
    # A unique ID (ex: com.mywebsite.com.myname).
    # If not provided, a UUID v4 ID is generated. 
    id = "local"
    
    # Contract source location and information
    repo {
        # The pubic github repository
        url = "https://github.com/ellcrys/contract-example" 
        # The github release tag or commit id (default: latest release)
        version = "bf699e05d21d77acce51ad6234e8d69a6d62132c"
        # The contract source code language
        language = "go"
        # Specify the ID of another cocoon to link to.
        # The contract will have the same privileges of the linked contract
        # and will become participate in load balancing requests coming into 
        # the linked cocoon code. 
        # Both contracts must be owned by same identity.
        link = ""
    }
    
    # Provide build information if the contract code requires it
    build {
        # The package manager to use (supported: glide, govendor)
        pkgMgr = "govendor"
    }
    
    # Resources to allocate to the contract's cocoon
    resources {
       set = "s2"
    }
    
    # Provide signatory information
    signatories {
        # The maximum number of signatories to accept
        max = 1
        # The number of signature required to approve a release
        threshold = 1
    }
    
    # Access control list stanza allows the contract
    # to allow or deny access to perform specific operations by other contracts.
    acl {
        # Allow all operations but deny the ability to create ledgers
        "*" = "allow deny-create-ledger deny-get"
    }
    
    # Firewall stanza determines the addresses the contract
    # can make outbound connections to.
    firewall {
        
        # If enabled, the contract will not be able to make outbound connections (Default: true)
        enabled = true
        
        # Firewall rules for outbound connections.
        # IP and DNS name is allowed. DNS name will be automatically resolved.
        rule = {
            destination = "google.com"
            destinationPort = "80"
            protocol = "tcp"    
        }
    }
    
    # Set environment variable. Use flags to 
    # enable special directives for individual variables.
    # @private flag will cause the value to never show up in any publicly accessible channel
    # @genRand32 generates a 32 byte random string 
    env {
        "MY_VAR"  = "some value 2"
        "MY_VAR2@unpin_once,private" = "yo"
        "SOME" = "THING"
    }
    env {
        "MY_VAR"  = "some value 2"
        "MY_VAR2@unpin_once,private" = "yo"
        "SOME" = "THING",
    }
}
