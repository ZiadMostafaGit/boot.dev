def check_ingredient_match(recipe, ingredients):
        
    output=[]
    counter=0
    size=len(recipe)
    for i in range(size):
       if recipe[i]!=ingredients[i]:
        counter+=1
        output.append(recipe[i])
        
    return (1-(counter/size))*100,output        
                




