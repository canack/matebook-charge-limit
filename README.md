# Matebook charge limit with CLI

This tool generated to changing battery charge threshold values via command line.

# Requirements
This tool working huawei matebook laptops with using huawei-wmi kernel driver.\
Usually the latest linux kernels come pre-installed.\
For more info, please visit : [Huawei-WMI](https://github.com/aymanbagabas/Huawei-WMI)


## Installation
Execute `install.sh` with your regular user.

```bash
  git clone https://github.com/canack/matebook-charge-limit.git
  cd matebook-charge-limit
  ./install.sh 
```
And now, you ready to go =)

## Usage/Examples

```bash
chargelimit # Prints how to use

# For example
chargelimit 70 # Maximum charge limit setting as 70
```


## License

[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)
