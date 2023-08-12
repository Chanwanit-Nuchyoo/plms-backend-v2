import sys
import builtins


def custom_input(prompts_iter, prompt=""):
    """
    Custom input function that fetches input values from an iterator.

    Args:
    - prompts_iter (iterator): An iterator over user inputs.
    - prompt (str, optional): A string that is printed as a prompt.

    Returns:
    - str: Next input from the iterator.
    """
    try:
        response = next(prompts_iter)
        print(prompt + response)
        return response
    except StopIteration:
        raise EOFError("Ran out of input values.")


def main():
    """
    Main function that overrides the built-in input, and then executes user's code.
    """
    # Check for sufficient arguments
    if len(sys.argv) < 3:
        print("Usage: script.py <user_code> <user_inputs>")
        sys.exit(1)

    user_code = sys.argv[1]
    user_inputs = sys.argv[2].split("\n")
    prompts_iter = iter(user_inputs)

    # Override the built-in input function
    builtins.input = lambda prompt="": custom_input(prompts_iter, prompt)

    # Execute the user's code
    exec(user_code)


if __name__ == "__main__":
    main()
