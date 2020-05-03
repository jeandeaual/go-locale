# GUI Test Application

## Building

First, install the `fyne` command:

```sh
$ GO111MODULE=off go get fyne.io/fyne/cmd/fyne
```

Then, build the example program:

* For iOS:

    ```sh
    $ fyne package -os ios -appID com.example.myapp -icon icon.png
    ```

* For Android:

    ```sh
    $ fyne package -os android -appID com.example.myapp -icon icon.png
    ```

* On desktop:

    ```sh
    $ go build
    ```

For more information, see [the Fyne documentation](https://fyne.io/develop/distribution).
