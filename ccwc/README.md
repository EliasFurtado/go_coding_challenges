# Go Word Count (ccwc)

Go Word Count (ccwc) is a simple command-line tool that provides functionality similar to the `wc` (word count) utility on Linux. It can count the number of lines, words, characters, and bytes in a given input.

## How to Use

### Prerequisites
- Make sure you have Go installed on your machine.

### Download and Build
1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/EliasFurtado/go_coding_challenges.git
    ```

2. Change into the project directory:

    ```bash
    cd ccwc
    ```

3. Run the following command to build the `ccwc` executable:

    ```bash
    make
    ```

### Run ccwc
After building the executable, you can run ccwc on a file or use it interactively:

1. To count lines, words, and bytes in a file:

    ```bash
    ./ccwc your_file.txt
    ```

2. To use ccwc interactively (input from stdin):

    ```bash
    echo "Hello, world!" | ./ccwc
    ```

### Available Flags

- `-c`: Count bytes.
- `-l`: Count lines.
- `-w`: Count words.
- `-m`: Count characters.

Use these flags to customize the output based on your specific requirements. For example:

```bash
./ccwc -l -w your_file.txt
