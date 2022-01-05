import os
import json
import matplotlib.pyplot as plt

def PoW(blockchain):
    dic = dict()
    directory = 'data/block_decentralization/' + blockchain.lower()
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                date = data['TimeStamp'][:10]
                if date not in dic.keys():
                    dic[date] = dict()
                miner = data['BlockMiner']
                if miner not in dic[date].keys():
                    dic[date][miner] = 0
                dic[date][miner] += 1
            except:
                print(f)
    dates = []
    nakamoto_coefficient = []
    for date in sorted(dic.keys()):
        dates.append(date)
        miners = sorted(list(dic[date].values()), reverse=True)
        miners = list(map(lambda x:x/sum(miners), miners))
        count = 0
        percent = 0
        for i in miners:
            percent += i
            count += 1
            if percent > 0.5:
                break
        nakamoto_coefficient.append(count)
    print(nakamoto_coefficient)

    fig = plt.figure()
    plt.plot(dates, nakamoto_coefficient, color ='b')
    plt.xlabel('date')
    plt.ylabel('nakamoto coefficient')
    plt.title(blockchain +' Block Production')
    # plt.show()
    plt.savefig('images/block_decentralization/'+ blockchain +'_nakamoto.png')

def PoS(blockchain):
    dic = dict()
    directory = 'data/num_actors/' + blockchain.lower()
    for filename in os.listdir(directory):
        file = os.path.join(directory, filename)
        if os.path.isfile(file):
            f = open(file)
            try:
                data = json.load(f)
                date = data['TimeStamp']
                if date not in dic.keys():
                    dic[date] = []
                validator = data['Account']
                stake = data['ActiveStake']
                dic[date].append(stake)
            except Exception as e:
                print(e)
    dates = []
    nakamoto_coefficient = []
    for date in sorted(dic.keys()):
        dates.append(date)
        validators = sorted(dic[date], reverse=True)
        validators = list(map(lambda x:x/sum(validators), validators))
        count = 0
        percent = 0
        for i in validators:
            percent += i
            count += 1
            if percent > 1/3:
                break
        nakamoto_coefficient.append(count)
    print(nakamoto_coefficient)

    fig = plt.figure()
    plt.plot(dates, nakamoto_coefficient, color ='b')
    plt.xlabel('Time')
    plt.ylabel('Nakamoto Coefficient')
    plt.title('Solana Validator Stake')
    # plt.show()
    plt.savefig('images/block_decentralization/solana_nakamoto.png')

PoS('Solana')