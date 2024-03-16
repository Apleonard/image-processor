# image-processor

## Contents

- [Prerequisite](#prerequisite)
- [How To Run](#how-to-run)
- [Hit it](#hit-it)

## Prerequisite

1. Go 1.21
2. gin 1.9.1
3. GoCV

## How to run

### 1. Install project dependencies

```sh
go mod tidy
```

### 2. Install GoCV

Read the documentation here: https://github.com/amacneil/dbmate

#### Quick Install

First, change directories to where you want to install GoCV, and then use git to clone the repository to your local machine like this:
cd $HOME/folder/with/your/src/
	git clone https://github.com/hybridgroup/gocv.git
Make sure to change `$HOME/folder/with/your/src/` to where you actually want to save the code.
Once you have cloned the repo, the following commands should do everything to download and install OpenCV 4.8.1 on Linux:
cd gocv
make install
If you need static opencv libraries
make install BUILD_SHARED_LIBS=OFF

If it works correctly, at the end of the entire process, the following message should be displayed:

    gocv version: 0.34.0
    gocv version: 0.35.0
    opencv lib version: 4.8.1

That's it, now you are ready to use GoCV.

Read the documentation here: https://github.com/hybridgroup/gocv

### 3. Run The Program

Program will automatically run on port 8080 if you didnt specify the port

```sh
go run cmd/main.go
```

## Hit it

Below is cURL to test it with your png files

```sh
curl --location 'localhost:8080/api/v1/img' \
--form 'file=@"/path/to/file/your_file.png"' \
--form 'width="180"' \
--form 'height="100"'
```
