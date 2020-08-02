package topology

import (
    "encoding/json"
    "io/ioutil"
)

type Topology struct {
    Producers []NodeConfig `json:"Producers"`
}

type NodeConfig struct {
    NodeName     string `json:"node,omitempty"`
    FriendlyName string `json:"friendly_name,omitempty"`
    Operator     string `json:"operator,omitempty"`
    HostAddress  string `json:"addr"`
    Port         int    `json:"port"`
    Valency      int    `json:"valency,omitempty"`
}

// reads the topology JSON file at the given file path. if
// parsing the file fails, an error will be returned.
func ReadTopologyFile(filePath string) (*Topology, error) {
    var topology Topology
    var data, err = ioutil.ReadFile(filePath)
    if err == nil {
        err = json.Unmarshal(data, &topology)
        if err == nil {
            return &topology, nil
        }
    }
    return nil, err
}

// writes the given topology to the given file.
func WriteTopologyFile(filePath string, topology *Topology) error {
    data, err := json.Marshal(topology)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filePath, data, 0644)
}
