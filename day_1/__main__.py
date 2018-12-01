import os
import sys


def help_message():
    return """Required Argument:
    part :: 1 :: Return the results of part one
    part :: 2 :: Return the results of part two

    USAGE:
    python -m day_1 `part`
    """

def read_frequencies(frequency_file):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    file_path = f'{dir_path}/{frequency_file}'
    with open(file_path, 'r') as f:
        frequency_list = f.read().splitlines()
    return frequency_list

def resulting_frequency(frequency_list):
    starting_frequency = [0]
    frequencies = starting_frequency + frequency_list
    return sum(int(frequency) for frequency in frequencies)


if __name__ == '__main__':
    try:
        if sys.argv[1] == '1':
            print(resulting_frequency(read_frequencies('frequencies.txt')))
        elif sys.argv[1] == '2':
            pass
        else:
            print(f"""Error: Invalid argument
            {help_message()}
            """)
    except IndexError:
        print(help_message())
