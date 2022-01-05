import requests
from bs4 import BeautifulSoup
import pandas as pd
import re
import time
import datetime
import json
import os

headers = {'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36'}

def EtherscanScraper(url):
    r = requests.get(url,headers=headers)
    soup = BeautifulSoup(r.content, 'html5lib')
    tags = soup.find_all('div', attrs = {'class':'col-md-9'}) 
    return soup.find_all('td')

def OwnershipScraper():

  dic = {'Rank': [], 'isContract': [], 'Address': [], 'Name Tag': [], 'Balance': [], 'Percentage': [], 'Txn Count': []}
  for k in range(1, 101, 1):
    page = EtherscanScraper('https://etherscan.io/accounts/'+str(k)+'?ps=100')
    l = len(page)
    for i in range(l//6):
      start = 6 * i
      dic['Rank'].append(str(page[start])[4:-5])
      addressTag = str(page[start+1])[4:-5]
      if addressTag.find('Contract') == -1:
        dic['isContract'].append(False)
      else:
        dic['isContract'].append(True)
      ind = addressTag.find('/address/')
      dic['Address'].append(addressTag[ind+9:ind+51])
      dic['Name Tag'].append(str(page[start+2])[4:-5])
      dic['Balance'].append("".join(map(str,re.findall(r'[.\d]', str(page[start + 3])[4:-5]))))
      dic['Percentage'].append(str(page[start+4])[4:-5].replace('%', ''))
      dic['Txn Count'].append(str(page[start+5])[4:-5].replace(',', ''))
    time.sleep(0.5)

  dataframe = pd.DataFrame(dic)
  currentTime = (datetime.datetime.utcnow().isoformat()).replace(':', '_')
  dataframe.to_csv('data/Ethereum_Ownership_' + currentTime + '.csv')
  print(dataframe)

def csvToJson():
  dataFrame = pd.read_csv('data/Ethereum_Ownership_2022-01-05T11_57_37.674301.csv')
  for _, row in dataFrame.iterrows():
    json_obj = {}
    json_obj['Address'] = row['Address']
    json_obj['Balance'] = row['Balance']
    json_obj['Blockchain'] = 'Ethereum'
    json_obj['TimeStamp'] = '2022-01-05T11_57_37.674301'.replace('_', ':')
    filename = 'data/ownership_decentralization/ethereum/'+ row['Address'] + '_' + '2022-01-05T11_57_37.674301' +'.json'
    with open(filename, 'w') as jsonFile:
      json.dump(json_obj, jsonFile)

# OwnershipScraper()
# csvToJson()