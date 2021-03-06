package servicemail

import (
	"encoding/json"
	"time"
)

// Exchange describes the RabbitMQ exchange name.
const Exchange = "servicemail"

// Message broker queue names.
const (
	RoutingQueue  = Exchange + ".routing"
	DeliveryQueue = Exchange + ".deliveries"
)

// Routing key prefixes.
const (
	IngressRoutingKey     = "ingress"
	DeliveryRoutingKey    = "delivery"
	RPCRoutingKey         = "rpc"
	RPCResponseRoutingKey = RPCRoutingKey + ".response"
)

// RPC call names.
const (
	RoutingCall = "routing"
)

// DefaultRPCTimeout represents the default RPC timeout.
const DefaultRPCTimeout = 5 * time.Second

// Marshal aliases. (Will make it easy to swap out serializers in future.)
var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
)
