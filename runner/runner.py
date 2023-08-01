import sys
import builtins


def main():
    # Load the user's code and the inputs from the command-line arguments
    user_code = sys.argv[1]
    user_inputs = sys.argv[2].split("\n")

    # This is the modified input function
    def input(prompt):
        # Get the next input from the list
        response = user_inputs.pop(0)

        # Print the prompt and the response
        print(prompt + response)

        # Return the response, so that the user's code can use it
        return response

    # Override the built-in input function
    builtins.input = input

    # Execute the user's code
    exec(user_code)


if __name__ == "__main__":
    main()
