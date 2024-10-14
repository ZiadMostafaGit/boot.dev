def remove_duplicates(spells):
    
    no_dup=set()

    for item in spells:
        no_dup.add(item)
    return list(no_dup)    





def remove_duplicates(lst):

        return set(lst)





def count_vowels(text):
       reslist=[]

       vawels=['a','A','e','E','o','O','i','I','u','U']
       for i in text:
         if i in vawels:
            reslist.append(i)
       
       return len(reslist),set(reslist)      
            
    
def find_missing_ids(first_ids, second_ids):

    lst=[]
    for i in first_ids:
        if i not in second_ids:
            lst.append(i)
    
    lst=set(lst)


    return list(lst)            
                        





def get_player_record(player_id):



    if player_id == 1:
        return {"name": "Slayer", "level": 128}
    if player_id == 2:
        return {"name": "Dorgoth", "level": 300}
    if player_id == 3:
        return {"name": "Saruman", "level": 4000}

    raise Exception("player id not found")




def handle_get_player_record(player_id):

    try:
        res=get_player_record(player_id)
    
    except IndexError:
            print("index is too high")    

    except Exception as e:

        print(e)








def purchase_item(price, gold_available):

    if price>gold_available:
        raise Exception("not enough gold")
    
    return gold_available-price









        
        





# Don't edit below this line


def get_player_record(player_id):
    if player_id < 0:
        raise Exception("negative ids not allowed")
    players = [
        {"name": "Slayer", "level": 128},
        {"name": "Dorgoth", "level": 300},
        {"name": "Saruman", "level": 4000},
    ]
    return players[player_id]






def process_transactions(purchase_orders):

    leftover=[]
    for order in purchase_orders:
        try:
            
           
            vals=list(order.values())
            res=purchase_item(vals[0],vals[1])
            leftover.append(res)
        except Exception as e:

            print(e)
    


    return leftover        






# Don't edit below this line


def main():
    print("Processing transactions...")
    leftovers = process_transactions(
        [
            {"price": 10.00, "gold_available": 125.00},
            {"price": 5.00, "gold_available": 2.00},
            {"price": 20.01, "gold_available": 5.20},
            {"price": 1.04, "gold_available": 254.00},
            {"price": 40.20, "gold_available": 6.00},
            {"price": 16.00, "gold_available": 235.01},
            {"price": 10.70, "gold_available": 10.70},
            {"price": 12.00, "gold_available": 2.30},
        ]
    )
    print("Transactions complete!")
    print("Leftover amounts for successful purchases:")
    for leftover in leftovers:
        print(f" * {leftover:.2f}")


def purchase_item(price, gold_available):
    if gold_available < price:
        raise Exception(f"{gold_available:.2f} is not enough for {price:.2f}")
    return gold_available - price


main()
