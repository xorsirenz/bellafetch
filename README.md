# bellafetch
<pre><code>minimal system information gathering tool<br></code></pre>


![](https://github.com/xorsirenz/bellafetch/blob/main/assets/ss-default-config.png?raw=true)

## Requirements:
- \>= [golang v1.22.0](https://go.dev/)

##  Installation:
### Build from source
```sh
# clone the repo
$ git clone https://github.com/xorsirenz/bellafetch.git 

# cd into the repo
$ cd bellafetch/

# build bellafetch
$ make

# move binary into your $PATH.. example:
$ sudo mv bellafetch /usr/local/bin
 ```
## Usage
 ```sh
 bellafetch
 ```
### Configuration

> [!IMPORTANT]
> Currently, bellafetch does not impliment JSON validation via JSON schema, so it is very important to follow the pattern shown in the default configuration below.
 
 Bellafetch will create a default config file in the location: 
 ```sh
 $ ~/.config/bellafetch/config
 ```

#### Default Config:
 ```json
{
  "Ascii": "none",
  "Modules": {
    "Cpu": true,
    "DiskSpace": true,
    "Gpu": true,
    "Host": true,
    "Kernel": true,
    "Memory": true,
    "Package": true,
    "PrettyName": true,
    "Shell": true,
    "Terminal": true,
    "Uptime": true,
    "WM": true
  }
}
```

### ASCII Art
#### Default ASCII art:
By default ASCII art is set to none/disabled. You can set it to default to match your distro.
Default will try and match your distros os-release "ID" or "ID_LIKE" value with a matching .txt filename in
[internal/utils/ascii](https://github.com/xorsirenz/bellafetch/tree/main/internal/utils/ascii).

#### Custom ASCII art:
To use a custom ASCII art of your choice, add a .txt file inside the directory 
[internal/utils/ascii](https://github.com/xorsirenz/bellafetch/tree/main/internal/utils/ascii), 
then specify the name of the file the in Ascii section of your config.
```json
{
    "Ascii": "custom",
    ...
}
```
