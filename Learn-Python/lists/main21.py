def validate_path(expected_path, character_path):
        character_name=character_path[0]
        size=len(expected_path)
        counter=0
        for i in range(size):
            if expected_path[i]!=character_path[i+1]:
                counter+=1



        return character_name,(1-(counter/size))*100
