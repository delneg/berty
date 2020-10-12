// file generated. see /config.
package config

import (
	"encoding/json"

	"berty.tech/berty/v2/go/pkg/bertytypes"
)

var Config bertytypes.Config

// FIXME: make it more nicely
func init() {
	err := json.Unmarshal([]byte("{\"berty\":{\"contacts\":{\"betabot\":{\"link\":\"https://berty.tech/id#key=CiBYAkJkmvcCZOl2hWuSK34arbzSpcpQGLowIvi7ZsEdyRIgMmKs-zHKksC74gjOfSj5puOAQQGWNhsC8o9gEtQ8zrQ\\u0026name=BetaBot\"},\"testbot\":{\"link\":\"https://berty.tech/id#key=CiCTp8uUnxfqlZaebiJEEt_zY-NUnHAWJl4jbGpB7rXsbxIg6_Grg-eLJ-GtqrdARnoTY35k7AI_g5AxqyFVe2elIaA\\u0026name=TestBot\"}}},\"p2p\":{\"rdvp\":[{\"maddr\":\"/ip4/51.159.21.214/tcp/4040/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p\"},{\"maddr\":\"/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p\"},{\"maddr\":\"/ip4/51.15.25.224/tcp/4040/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ\"},{\"maddr\":\"/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ\"},{\"maddr\":\"/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWRpyQpZtUmY5ZktEMgzuhNoWC1C9zokDjLVahNMy3g48u\"}]}}"), &Config)
	if err != nil {
		panic(err)
	}
}