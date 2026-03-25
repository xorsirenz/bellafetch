# bellafetch
<pre><code>minimal system information gathering tool<br></code></pre>


![](https://github.com/xorsirenz/bellafetch/blob/main/ss.png?raw=true)

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
  "Modules": [
    {
      "Host": true,
      "PrettyName": true,
      "Kernel": true,
      "Uptime": true,
      "Packages": true,
      "Shell": true,
      "Terminal": true,
      "WM": false,
      "Cpu": true,
      "Gpu": true,
      "DiskSpace": true,
      "Memory": true
    }
  ]
}
```
