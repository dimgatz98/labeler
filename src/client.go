package labeler

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	labeler "github.com/dimgatz98/labeler/src/server"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := labeler.NewLabelerServiceClient(conn)

	data := labeler.Label{
		Node:  "k3d-gputest-server-0",
		Label: "wooohoooo: 1234",
	}

	response, err := c.LabelNode(context.Background(), &data)
	if err != nil {
		log.Fatalf("Error when calling LabelNode: %v", err)
	}

	log.Printf("Response from server: %v", response)
}
