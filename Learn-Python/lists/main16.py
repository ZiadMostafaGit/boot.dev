def reverse_array(items):
    linght=len(items)-1
    size=int(len(items)/2)
    for i in range(size):
      temp=items[i]
      items[i]=items[linght-i]
      items[linght-i]=temp

    return items
