def number_sum(n):
    return (n*(n+1))/2














def find_min(nums):

    min=float("inf")
    for i in nums:
        if i<min:
            min=i



    return min



def remove_nonints(nums):
   new_list=[]
   for num in nums:
       if type(num)==int:
           new_list.append(num)


   return new_list






def factorial(num):
        if num<=1:
            return 1

        return num*factorial(num-1)













def area_sum(rectangles):

    rec=0 
    for area in rectangles:
        res=1 
        for item in area:

          res*=area[item] 

        rec+=res



    return rec 





def fizzbuzz(start, end):


    for i in range(start,end):
        if i %3==0 and i%5==0:
            print("fizzbuzz")


        if i%3==0:
            print("fizz")


        if i%5==0:
            print("buzz")


        else:
            print(i)









    






# Don't Touch Below This Line





def divide_list(nums, divisor):
    




    for num in range(nums):

        nums[num]=nums[num]/divisor

    return nums




def join_strings(strings):
    

        return ",".join(strings)
