package main

import (
    "encoding/json"
    "log"

    "github.com/google/uuid"
    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
    // Unmarshal the message body as an loosely-typed map.
    var body map[string]any
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }

    result := make(map[string]any)
    // Update the message type to return back.
    result["type"] = "generate_ok"
    result["id"] = uuid.New().String()

    // Echo the original message back with the updated message type.
    return n.Reply(msg, result)
	})

	if err := n.Run(); err != nil {
    log.Fatal(err)
}
}
