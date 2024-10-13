def filter_messages(messages):
    new_list=[]
    indexlist=[]
    for message in messages:
        wordlist=list(message.split())
        counter=0
        while "dang" in wordlist:
            wordlist.remove("dang")
            counter+=1



        new_list.append(" ".join(wordlist))
        #new_list.append(str(wordlist))
        indexlist.append(counter)


    return new_list,indexlist




test=["hi dang dang it in row"," ","dang her","iam ziad"]

newlist,index_list=filter_messages(test)
print(newlist)



