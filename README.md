# DistributedAI

This is the distributed AI project @ zju.

We will focus on the research of:
1. Distributed Machine Learning
2. Blockchain

## Installation

### Docker installation 

We assume you have a command line interface (CLI) in your OS 
(bash, zsh, cygwin, git-bash, power-shell etc.). We assume this CLI sets 
 the variable `$(pwd)` to the current directory. If it doesn't replace
 all mentions of `$(pwd)` with the current directory you are in. 

#### Install Docker

Go to the [docker webpage](https://www.docker.com/) and follow the instruction for your platform.

#### Download Image

Next you can download the docker image:

    docker pull lzhou1110/distributed_ai_zju
    
#### Get Git Repository

You can use the git installation in the docker container to get the repository:

    docker run -v "$(pwd)":/home/zju/work lzhou1110/distributed_ai_zju git clone https://github.com/lzhou1110/DistributedAI.git

Note: this will create a new `DistributedAI` directory in your current directory.

#### Change into directory

    cd DistributedAI

Note: you need to be in the `DistributedAI` directory every time you want to run/update the book.

#### Run Notebook

    docker run -it --rm -p 8888:8888 -v "$(pwd)":/home/zju/work lzhou1110/distributed_ai_zju 

You are now ready to visit the jupyter notebook at http://localhost:8888
 
## Usage

Once installed you can always run your notebook server by first changing
into your local `DistributedAI` directory, and then executing:

    docker run -it --rm -p 8888:8888 -v "$(pwd)":/home/zju/work lzhou1110/distributed_ai_zju 
    
This is **assuming that your docker daemon is running** and that you are
**in the `DistributedAI` directory**. How to run the docker daemon
depends on your system.


For markdown tutorial, refer to: <code>https://guides.github.com/features/mastering-markdown/</code>

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## Mendeley Group
We also have a Mendeley group: <code>mendeley.com/community/distributedaizju</code>

## History

The tensorflow mnist examples were from: <code>https://github.com/ianlewis/tensorflow-examples</code>

## Credits

This project is led by (in alphabetic order):

* [Chao Wu]() @wuchaozju
* [Jun Xiao](junx@cs.zju.edu.cn) junx@cs.zju.edu.cn

## License
Apache License 2.0