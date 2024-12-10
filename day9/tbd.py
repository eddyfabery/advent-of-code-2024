def prep_data():
    with open("test-input.txt", "r") as f:
        data = f.read().splitlines()
        data = [int(x) for x in data[0]]
    return data


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


start = 0
end = len(expanded_data)-1

while start <= end:
    a = expanded_data[start]
    b = expanded_data[end]

    if b != ".":
        # Scan backwards and find the length of file block to swap with
        j = end
        while start < j and expanded_data[j] == b:
            print("INNER ITERATION for j", j, expanded_data[j:end+1])
            j -= 1

        block_len = end-j
        print("LENGTH OF BLOCK TO MOVE", block_len)

        # Scan forward and find the length of file available freespace
        k = start
        avail_space_len = 0
        while k <= j:
            print(f"K: {k}, Space: {avail_space_len}, Block: {block_len}")
            if expanded_data[end] == 7:
                breakpoint()

            if expanded_data[k] != '.':
                k+=1
                avail_space_len = 0
                continue
            elif expanded_data[k] == "." and avail_space_len < block_len:
                avail_space_len += 1
                k+=1
                continue
            elif avail_space_len == block_len:
                print("+++++SWAP++++++++++++++++")
                ix = k-avail_space_len
                print("SWAPPING:", ''.join(expanded_data[ix: k]),  ' and ', ''.join(expanded_data[j+1:end+1]))
                expanded_data[ix: k], expanded_data[j+1:end+1] = expanded_data[j+1:end+1], expanded_data[ix: k]
                print("OUTCOME", ''.join(expanded_data))
                end = j
                break
        if k >= j:
            print("+++++UNABLE TO MOVE    !!!!!!!", ''.join(expanded_data[j+1:end+1]))
            end = end-block_len
    else:
        end -= 1
    breakpoint()
print(''.join(expanded_data))
