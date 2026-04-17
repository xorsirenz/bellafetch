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
#### Config location: 
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
