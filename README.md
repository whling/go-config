# go-config
go应用通用配置处理

## 使用方式
go get -u github.com/whling/go-config



func main() {
    toml, err := config.LoadToml("./test.toml")
    if (err != nil) {
      painc()
    }
    
    tomlValue := &TomlValue{toml}
    b := tomlValue.Bool("startHttp", true)
    if b {
       fmt.Println("startHttp")
    }

}
