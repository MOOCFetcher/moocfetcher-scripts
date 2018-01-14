### Setup EC2 Instance

#### Install dependencies
```
sudo apt-get install git virtualenv aria2
```

Other dependencies include:

* moreutils
* jq
* pup

#### Setup virtualenv w/ Python 3
```
virtual-env -python=python3
```

Make sure you use `pip3` to install Python related software.

#### Install coursera-dl
[Setup coursera-dl from source](https://github.com/coursera-dl/coursera-dl#alternative-installation-method-for-unix-systems) inside a virtualenv.


#### Install awscli and other python deps

Run this inside the virtualenv:

```
pip install awscli csvkit
```

