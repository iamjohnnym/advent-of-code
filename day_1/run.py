import os
import sys


def help_message():
    return """Required Argument:
    part :: 1 :: Return the results of part one
    part :: 2 :: Return the results of part two

    USAGE:
    python -m day_1.run `part`
    """

def read_frequencies(frequency_file):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    file_path = f'{dir_path}/{frequency_file}'
    with open(file_path, 'r') as f:
        frequency_list = f.read().splitlines()
    return frequency_list

def get_frequencies(frequency_list):
    return frequency_list

def resulting_frequency(frequency_list):
    frequencies = get_frequencies(frequency_list)
    return sum(int(frequency) for frequency in frequencies)

def first_repeated_frequency(frequency_list):
    frequencies = get_frequencies(frequency_list)
    frequency_results = {0}
    last_result = 0
    while True:
        for frequency in frequencies:
            last_result += int(frequency)
            if not last_result in frequency_results:
                frequency_results.add(last_result)
            else:
                return last_result


if __name__ == '__main__':
    try:
        if sys.argv[1] == '1':
            print(resulting_frequency(read_frequencies('frequencies.txt')))
        elif sys.argv[1] == '2':
            print(first_repeated_frequency(read_frequencies('frequencies.txt')))
        else:
            print(f"""Error: Invalid argument
            {help_message()}
            """)
    except IndexError:
        print(help_message())
