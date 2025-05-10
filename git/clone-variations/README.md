# Clone Variations

The script `clone-variations.sh` demonstrates different ways to clone a git repository, with the goal of finding the smallest resulting repository size thereby reducing the time it takes to clone a repository.

## Usage

```bash
./clone-variations.sh https://repo-url
```

## Requirements

A Unix-like environment with the following tools installed:

- `Bash`
- `Git`
- GNU utils
  - `parallel`
  - `sort`
  - `du`

The script will take a loongish time to run, go listen to some music or something.

## Sample Run

Cloning the gorm repo (9th May, 2025).

```
./clone-variations.sh https://github.com/go-gorm/gorm.git

Computers / CPU threads / Max jobs to run
1:local / 8 / 8

Computer:jobs running/jobs completed/%of started jobs/Average seconds to complete
local:8/0/100%/0.0s Cloning into bare repository 'clones/blobless-shallow-clone'...
local:8/1/100%/2.0s Cloning into 'clones/shallow-clone'...
local:7/2/100%/1.0s Cloning into bare repository 'clones/treeless-bare-clone'...
local:6/3/100%/0.7s Cloning into bare repository 'clones/bare-blobless-clone'...
local:5/4/100%/0.8s Cloning into bare repository 'clones/treeless-single-branch-clone'...
local:4/5/100%/0.6s Cloning into 'clones/plain-clone'...
local:3/6/100%/0.5s Cloning into bare repository 'clones/blobless-single-branch-clone'...
local:2/7/100%/0.4s Cloning into bare repository 'clones/bare-clone'...
local:1/8/100%/0.6s Cloning into 'clones/single-branch-clone'...
local:0/9/100%/0.6s 

112K    clones/blobless-shallow-clone
892K    clones/treeless-bare-clone
1.3M    clones/treeless-single-branch-clone
1.6M    clones/bare-blobless-clone
1.8M    clones/shallow-clone
2.8M    clones/blobless-single-branch-clone
5.8M    clones/bare-clone
7.0M    clones/plain-clone
7.5M    clones/single-branch-clone
```
