def get_heroes():
    heroes = [
        "Glorfindel",
        2093,
        True,
        "Gandalf",
        1054,
        False,
        "Gimli",
        389,
        False,
        "Aragorn",
        87,
        False,
    ]

    newlist=[]
    size=len(heroes)
    for i in range(0,size,3):
        newlist.append(tuple(heroes[i:i+3]))






    return newlist 


# Don't touch below this line


def test():
    heroes = get_heroes()
    for hero in heroes:
        print(f"name: {hero[0]}, age: {hero[1]}, is_elf: {hero[2]}")


def main():
    test()


main()

