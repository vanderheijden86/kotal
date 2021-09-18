package polkadot

import (
	"fmt"
	"os"

	polkadotv1alpha1 "github.com/kotalco/kotal/apis/polkadot/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
)

// PolkadotClient is parity Polkadot client
// https://github.com/paritytech/polkadot
type PolkadotClient struct {
	node *polkadotv1alpha1.Node
}

// Images
const (
	// EnvPolkadotImage is the environment variable used for polkadot client image
	EnvPolkadotImage = "POLKADOT_IMAGE"
	// DefaultPolkadotImage is the default polkadot client image
	DefaultPolkadotImage = "parity/polkadot:v0.9.9-1"
	//  PolkadotHomeDir is go ipfs image home dir
	PolkadotHomeDir = "/polkadot"
)

// Image returns go-ipfs image
func (c *PolkadotClient) Image() string {
	if os.Getenv(EnvPolkadotImage) == "" {
		return DefaultPolkadotImage
	}
	return os.Getenv(EnvPolkadotImage)
}

// Command is go-ipfs entrypoint
func (c *PolkadotClient) Command() []string {
	return nil
}

// Args returns go-ipfs args
func (c *PolkadotClient) Args() (args []string) {

	node := c.node

	args = append(args, PolkadotArgBasePath, shared.PathData(c.HomeDir()))
	args = append(args, PolkadotArgChain, node.Spec.Network)
	args = append(args, PolkadotArgSync, string(node.Spec.SyncMode))
	args = append(args, PolkadotArgLogging, string(node.Spec.Logging))

	if node.Spec.RPC {
		args = append(args, PolkadotArgRPCExternal)
		args = append(args, PolkadotArgRPCPort, fmt.Sprintf("%d", node.Spec.RPCPort))
	}

	if node.Spec.WS {
		args = append(args, PolkadotArgWSExternal)
	}

	return
}

func (c *PolkadotClient) HomeDir() string {
	return PolkadotHomeDir
}