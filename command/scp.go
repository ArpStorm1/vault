package command

import (
	"strings"
)

type SCPCommand struct {
	// Because we are using sshpass to implement the scp functionality, we can compose the SSHCommand
	*SSHCommand

	sourcePath      string
	destinationPath string
	isRecursive     string
}

func (c *SCPCommand) Help() string {
	helpText := `
Usage: vault scp [options] -source=source-path -dest=dest-path [-recursive] [ssh options]

  Establishes an SCP connection with the target machine on top of SSH.

  This command uses one of the SSH secrets engines to authenticate and
  automatically establish an SCP connection to a host. This operation requires
  that the SSH secrets engine is mounted and configured.

  SCP using the OTP mode (requires sshpass for full automation):
	
	Copy file from current host to remote target:

      $ vault scp -mode=otp -role=my-role -source=path-to-file -dest=user@1.2.3.4:/path-to-put-the-file

	Copy file from remote target to current host:

	  $ vault scp -mode=otp -role=my-role -source=user@1.2.3.4:/path-to-file -dest=path-to-put-the-file

	Copy a directory from current host to remote target:
	
	  $ vault scp -mode=otp -role=my-role -source=path-to-dir -dest=user@1.2.3.4:/path-to-put-the-file -recursive

  SSH using the CA mode:

      $ vault scp -mode=ca -role=my-role -source=path-to-dir -dest=user@1.2.3.4:/path-to-put-the-file -recursive

  SSH using CA mode with host key verification:

      $ vault ssh \
          -mode=ca \
          -role=my-role \
          -host-key-mount-point=host-signer \
          -host-key-hostnames=example.com \
          -source=path-to-dir 
		  -dest=user@1.2.3.4:/path-to-put-the-file 
		  -recursive


  For the full list of options and arguments, please see the documentation.

` + c.Flags().Help()

	return strings.TrimSpace(helpText)
}
