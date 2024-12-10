def prep_data():
    with open("input.txt", "r") as f:
        data = f.read().splitlines()
        data = [int(x) for x in data[0]]
    return data


def main():
    data = prep_data()

    expanded_data = []
    current_id = 0
    file_block = True

    for i in data:
        if file_block:
            expanded_data += [str(current_id)] * i
            current_id += 1
        else:
            expanded_data += ["."] * i

        file_block = not file_block

    expanded_data = swap_spaces(expanded_data)
    checksum = get_checksum(expanded_data)
    print(f"Checksum: {checksum}")

def swap_spaces(expanded_data):
    start = 0
    end = len(expanded_data)-1
    while start <= end:
        a = expanded_data[start]
        b = expanded_data[end]

        if a == "." and b != ".":
            expanded_data[start], expanded_data[end] = expanded_data[end], expanded_data[start]
            start += 1
            end -=1
        elif a != "." and b != ".":
            start +=1
        elif a == "." and b == ".":
            end -= 1
        elif a != "." and b == ".":
            start += 1
            end -= 1

    return expanded_data

def get_checksum(data):
    sum = 0
    for i in range(len(data)):
        if data[i] != ".":
            sum += i * int(data[i])

    return sum


if __name__ == "__main__":
    main()
