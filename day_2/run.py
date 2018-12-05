import os
import sys


def help_message():
    return """Required Argument:
    part :: 1 :: Return the results of part one
    part :: 2 :: Return the results of part two

    USAGE:
    python -m day_2.run `part`
    """

def read_file(filename):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    file_path = f'{dir_path}/{filename}'
    with open(file_path, 'r') as f:
        lines = f.read().splitlines()
    return lines

def checksum_count(checksum):
    """
    >>> checksum_count('abcdef')
    {'two': 0, 'three': 0}
    >>> checksum_count('bababc')
    {'two': 1, 'three': 1}
    >>> checksum_count('abbcde')
    {'two': 1, 'three': 0}
    >>> checksum_count('abcccd')
    {'two': 0, 'three': 1}
    >>> checksum_count('aabcdd')
    {'two': 1, 'three': 0}
    >>> checksum_count('abcdee')
    {'two': 1, 'three': 0}
    >>> checksum_count('ababab')
    {'two': 0, 'three': 1}
    """

    boxes = { 'two': 0, 'three': 0 }
    letter_checked = set()
    for letter in checksum:
        if letter not in letter_checked:
            letter_checked.add(letter)
            total = checksum.count(letter)
            if total == 2 and boxes['two'] < 1:
                boxes['two'] += 1
            elif total == 3 and boxes['three'] < 1:
                boxes['three'] += 1
    return boxes

def get_checksum_from_file(scans):
    boxes = { 'two': 0, 'three': 0 }
    for checksum in scans:
        checksum_result = checksum_count(checksum)
        boxes['two'] += checksum_result['two']
        boxes['three'] += checksum_result['three']
    return boxes

def calculate_matches(checksum_file):
    total_matches = get_checksum_from_file(checksum_file)
    return total_matches['two'] * total_matches['three']

if __name__ == '__main__':
    try:
        if sys.argv[1] == '1':
            print(calculate_matches(read_file('challenge.txt')))
        elif sys.argv[1] == '2':
            pass
        elif sys.argv[1] == 't':
            import doctest
            doctest.testmod()
        else:
            print(f"""Error: Invalid argument
            {help_message()}
            """)
    except IndexError:
        print(help_message())
