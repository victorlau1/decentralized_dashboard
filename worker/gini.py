import json
import os
import numpy as np
import matplotlib.pyplot as plt

def plotGini(image_path, sorted_list, cumulative_list, gini_coefficient):
    x_axis = []
    for i in range(len(sorted_list)):
        x_axis.append(i)
    x_axis = list(map(lambda x: x/x_axis[-1] * 100, x_axis))
    y_axis = list(map(lambda x: x/cumulative_list[-1] * 100, cumulative_list))
    fig = plt.figure()
 
    # creating the lorenz curve
    plt.bar(x_axis, y_axis, color ='b', label='lorenz curve')
    x = np.linspace(0, 100, 100)
    plt.plot(x, x, '-r', label='Ideal')
    # plt.xlabel('')
    # plt.ylabel('')
    vector = image_path.split('/')[1] + ' ' + image_path.split('/')[2].split('.')[0]
    plt.title(vector + " gini = " + str(gini_coefficient))
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

def devDecentralization(directory, image_path):
    dic = dict()
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        # checking if it is a file
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                i = data['Committer']
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


blockDecentralization('data/block_decentralization/bitcoin', 'images/block_decentralization/bitcoin.png')
blockDecentralization('data/block_decentralization/ethereum', 'images/block_decentralization/ethereum.png')
devDecentralization('data/github/bitcoin', 'images/dev_decentralization/bitcoin.png')
devDecentralization('data/github/erigon', 'images/dev_decentralization/erigon.png')
devDecentralization('data/github/geth', 'images/dev_decentralization/geth.png')
devDecentralization('data/github/nethermind', 'images/dev_decentralization/nethermind.png')
devDecentralization('data/github/openethereum', 'images/dev_decentralization/openethereum.png')
devDecentralization('data/github/solana', 'images/dev_decentralization/solana.png')
ownershipDecentralization('data/ownership_decentralization/bitcoin', 'images/ownership_decentralization/bitcoin.png')
ownershipDecentralization('data/ownership_decentralization/solana', 'images/ownership_decentralization/solana.png')