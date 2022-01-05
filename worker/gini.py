import json
import os
import numpy as np
import matplotlib.pyplot as plt
import math

def plotGini(image_path, sorted_list, cumulative_list, gini_coefficient):
    x_axis = []
    for i in range(len(sorted_list)):
        x_axis.append(i)
    x_axis = list(map(lambda x: x/x_axis[-1] * 100, x_axis))
    y_axis = list(map(lambda x: x/cumulative_list[-1] * 100, cumulative_list))
    fig = plt.figure()
 
    # creating the lorenz curve
    plt.bar(x_axis, y_axis, color ='royalblue', label='Lorenz Curve')
    x = np.linspace(0, 100, 100)
    plt.plot(x, x, color='dimgray', label='Line of Ideal Decentralization')
    plt.xlabel('Cumulative Percentage of Addresses')
    plt.ylabel('Cumulative Percentage of Ownership')
    plt.suptitle('Ethereum Onwership Decentralization')
    plt.title('gini = ' + str(gini_coefficient)[:4], fontsize=10)
    # vector = image_path.split('/')[1] + ' ' + image_path.split('/')[2].split('.')[0]
    # plt.title(vector + " gini = " + str(gini_coefficient))
    plt.legend()
    plt.savefig(image_path)
    # plt.show()

def gini(list_of_values, image_path):
    sorted_list = sorted(list_of_values)
    cumulative_list = []
    height, area = 0, 0
    for value in sorted_list:
        height += value
        cumulative_list.append(height)
        area += height - value / 2.
    fair_area = height * len(list_of_values) / 2.
    gini_coefficient = (fair_area - area) / fair_area
    plotGini(image_path, sorted_list, cumulative_list, gini_coefficient)
    return gini_coefficient


def blockDecentralization(directory, image_path):
    dic = dict()
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        # checking if it is a file
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                i = data['BlockMiner']
                if i not in dic.keys():
                    dic[i] = 0
                dic[i] += 1
            except:
                print(f)
    print(gini(list(dic.values()), image_path))

    #hoover index
    # hashrate = list(dic.values())
    # mean = sum(hashrate) / len(hashrate)
    # res = 0
    # for node in hashrate:
    #     res += abs(mean - node)
    # res = res / sum(hashrate)
    # res /= 2
    # print(res)

    # theirl T index
    # hashrate = list(dic.values())
    # mean = sum(hashrate) / len(hashrate)
    # res = 0
    # for node in hashrate:
    #     res += (node / mean) * math.log(node/mean)
    # res = res / len(hashrate)
    # print(res)

    # theirl L index
    # hashrate = list(dic.values())
    # mean = sum(hashrate) / len(hashrate)
    # res = 0
    # for node in hashrate:
    #     res += math.log(mean/node)
    # res = res / len(hashrate)
    # print(res)

    # from the researchgate paper
    # arr = sorted(list(dic.values()))
    # arr = list(map(lambda x: x/sum(arr) * 100, arr))
    # dq = 0
    # for i in range(len(arr)):
    #     dq += arr[i] / (i+1) * 2
    # print(dq)

def devDecentralization(directory, image_path):
    dic = dict()
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        # checking if it is a file
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                i = data['Author']
                if i not in dic.keys():
                    dic[i] = 0
                dic[i] += 1
            except:
                print(f)
                # break
    print(sum(list(dic.values())))
    print(gini(list(dic.values()), image_path))

def ownershipDecentralization(directory, image_path):
    dic = []
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        # checking if it is a file
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                i = data['Balance']
                dic.append(i)
            except:
                print(f)
                # break
    
    print(len(dic))
    print(gini(dic, image_path))


# blockDecentralization('data/block_decentralization/bitcoin', 'images/block_decentralization/bitcoin.png')
# blockDecentralization('data/block_decentralization/ethereum', 'images/block_decentralization/ethereum.png')
# devDecentralization('data/dev_decentralization/bitcoin', 'images/dev_decentralization/bitcoin.png')
# devDecentralization('data/dev_decentralization/erigon', 'images/dev_decentralization/erigon.png')
# devDecentralization('data/dev_decentralization/geth', 'images/dev_decentralization/geth.png')
# devDecentralization('data/dev_decentralization/nethermind', 'images/dev_decentralization/nethermind.png')
# devDecentralization('data/dev_decentralization/openethereum', 'images/dev_decentralization/openethereum.png')
# devDecentralization('data/dev_decentralization/solana', 'images/dev_decentralization/solana.png')
# ownershipDecentralization('data/ownership_decentralization/bitcoin', 'images/ownership_decentralization/bitcoin.png')
# ownershipDecentralization('data/ownership_decentralization/ethereum', 'images/ownership_decentralization/ethereum.png')
# ownershipDecentralization('data/ownership_decentralization/solana', 'images/ownership_decentralization/solana.png')