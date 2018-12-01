import os


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
    print(resulting_frequency(read_frequencies('frequencies.txt')))

