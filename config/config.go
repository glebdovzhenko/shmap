package config


type ShmapCfg struct{
    Version []int
    Name string
}


func Default() (*ShmapCfg) {
    result := ShmapCfg{}
    result.Version = []int{0, 0, 1}
    result.Name = "SHMAP!"
    return &result
}
