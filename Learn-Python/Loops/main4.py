def calculate_flurry_crit(num_attacks, base_damage):
    totalDamage=0
    for i in range(num_attacks-1):
        totalDamage+=base_damage*2
    

    totalDamage+=base_damage*4
    return totalDamage

