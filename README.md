![image](https://github.com/deleonn/timezone_tracker/assets/15114660/de998b84-25d4-4ca4-9c4c-e5ffd6cd41bd)

# How to run

1. Build with `go build`.
2. Execute with nouhp `nohup ./timezone_tracker &`.
    1. nohup ignores the "hangup" signal from the terminal, allowing you to run the process independently. 
    2. The & at the end allows you to run the process in the background.
3. (Optional) Create a symlink (so you don't have to `cd` into the project to run the app) `ln -s /path/to/timezone_tracker /usr/local/bin/timezone_tracker`.
