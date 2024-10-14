# def get_character_record(name, server, level, rank):

#     char_record={
#         "name":name,
#         "server":server,
#         "level":level,
#         "rank":rank,
#         "id":name+"#"+server


#     }


#     return char_record



def get_character_record(name, server, level, rank):
    return {
        "name": name,
        "server": server,
        "level": level,
        "rank": rank,
        "id": f"{name}#{server}",
    }




def count_enemies(enemy_names):
    enemies_dict = {}
    for enemy_name in enemy_names:
        if enemy_name not in enemies_dict:
            enemies_dict[enemy_name]=1
        else:

            enemies_dict[enemy_name] += 1
    return enemies_dict




def get_most_common_enemy(enemies_dict):
    if enemies_dict!={}:

        max=float("-inf")
        resulte=""
        for enemie in enemies_dict:
            count=enemies_dict[enemie]
            if count >max:
                resulte=enemie
                max=count

        return resulte

    return None



def get_quest_status(progress):
    # Safely access nested dictionaries using chained key access
    return progress.get("entity", {}).get("character", {}).get("quests", {}).get("Dragon_Slayer", {}).get("status", None)













progress={
    "entity": {
        "character": {
            "name": "Kaladin",
            "quests": {
                "bridge_run": {
                    "status": "In Progress",
                },
                "talk_to_syl": {
                    "status": "Completed",
                },
            },
        }
    }
}    



get_quest_status(progress)



def merge(dict1, dict2):
    

    newdic={}
   
    newdic=dict1.copy()
    newdic.update(dict2)


    return newdic    







def total_score(score_dict):
  
    score=0
    for item in score_dict:
        score+=score_dict[item]


    return score    



def calculate_total(items_purchased, pinned_list):
    item_prices = {
        "health_potion": 10.00,
        "mana_potion": 12.00,
        "gold_dust": 5.00,
        "dwarven_ale": 8.00,
        "enchanted_scroll": 25.00,
        "ice_cold_milk": 50.00,
        "herbs": 7.00,
        "crystal_shard": 20.00,
        "magic_ring": 100.00,
        "mystic_amulet": 150.00,
    }

    # Don't touch above this line


    total=0
    items={}
    first_list=[]
    for item in pinned_list:
        if item not in items_purchased:
            first_list.append(item)
            
    for item in items_purchased:
        items[item]=item_prices[item]

    for item in items:
        total+=items[item]
    


    return first_list,items,total    

        

#
"""

Outputs
The function should return 3 values in this order:

unpurchased_items: A list of all the item names in pinned_list that weren't found in items_purchased in order.
receipt: A dictionary containing all the items the player purchased, even stuff not on their pinned_list. The keys are the item names and the values are their respective prices from the item_prices dictionary.
total: The total cost of all the items that were purchased.



"""#    






    
