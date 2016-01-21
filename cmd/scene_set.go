package cmd

import "fmt"

type SceneSet struct {
	SceneID   string
	SceneName string
}

func (c *SceneSet) FriendlyString() string {
	return fmt.Sprintf("Set scene \"%s\" [%s]", c.SceneName, c.SceneID)
}
func (c *SceneSet) String() string {
	return "cmd.SceneSet"
}
