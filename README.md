Yandex.Tracker API in Go
===============
This is the original Yandex.Tracker library for Go.

## Install

###
    go get -u github.com/dvsnin/yandex-tracker-go

## Example

### Getting ticket and description

```golang
import (
    "fmt"
	
    "github.com/dvsnin/yandex-tracker-go"
)

func main() {
    client := tracker.New("YOUR YANDEX.TRACKER TOKEN", "YOUR YANDEX ORG_ID")
    ticket, err := client.GetTicket("TICKET KEY")
    if err != nil {
    	fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("%s\n", ticket.Description())
}
```

### Edit ticket fields

```golang
import (
    "fmt"

    "github.com/dvsnin/yandex-tracker-go"
)

func main() {
    client := tracker.New("YOUR YANDEX.TRACKER TOKEN", "YOUR YANDEX ORG_ID")
    ticket, err := client.PatchTicket("TICKET KEY", map[string]string{"TICKET FIELD": "NEW VALUE"})
    if err != nil {
    	fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("%s\n", ticket.Description())
}
```

## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/dvsnin/yandex-tracker-go/blob/master/LICENSE) file for details
