from main import *

run_cases = [
    (
        (
            [
                "health_potion",
                "mana_potion",
                "gold_dust",
                "herbs",
                "crystal_shard",
                "dwarven_ale",
            ],
            [
                "health_potion",
                "mana_potion",
                "ice_cold_milk",
                "gold_dust",
                "herbs",
                "crystal_shard",
                "magic_ring",
                "dwarven_ale",
                "mystic_amulet",
            ],
        ),
        (
            ["ice_cold_milk", "magic_ring", "mystic_amulet"],
            {
                "health_potion": 10.00,
                "mana_potion": 12.00,
                "gold_dust": 5.00,
                "herbs": 7.00,
                "crystal_shard": 20.00,
                "dwarven_ale": 8.00,
            },
            62.00,
        ),
    ),
]

submit_cases = run_cases + [
    (
        (
            ["health_potion", "gold_dust", "herbs", "crystal_shard"],
            [
                "health_potion",
                "mana_potion",
                "gold_dust",
                "ice_cold_milk",
                "herbs",
                "magic_ring",
                "crystal_shard",
                "mystic_amulet",
            ],
        ),
        (
            ["mana_potion", "ice_cold_milk", "magic_ring", "mystic_amulet"],
            {
                "health_potion": 10.00,
                "gold_dust": 5.00,
                "herbs": 7.00,
                "crystal_shard": 20.00,
            },
            42.00,
        ),
    ),
    (
        (
            [
                "health_potion",
                "mana_potion",
                "gold_dust",
                "ice_cold_milk",
                "herbs",
                "magic_ring",
                "crystal_shard",
                "mystic_amulet",
            ],
            ["health_potion", "gold_dust", "herbs", "crystal_shard"],
        ),
        (
            [],
            {
                "health_potion": 10.00,
                "mana_potion": 12.00,
                "gold_dust": 5.00,
                "ice_cold_milk": 50.00,
                "herbs": 7.00,
                "magic_ring": 100.00,
                "crystal_shard": 20.00,
                "mystic_amulet": 150.00,
            },
            354.00,
        ),
    ),
]


def test(input1, expected_output):
    print("---------------------------------")
    print(f"Inputs: {input1}")
    print(f"Expecting: {expected_output}")
    try:
        unpurchased_items, receipt, total_cost = calculate_total(*input1)
    except Exception as e:
        print(f"Error: {e}")
        print("Fail")
        return False
    result = (unpurchased_items, receipt, total_cost)
    print(f"Actual: {result}")
    if result == expected_output:
        print("Pass")
        return True
    print("Fail")
    return False


def main():
    passed = 0
    failed = 0
    for test_case in test_cases:
        correct = test(*test_case)
        if correct:
            passed += 1
        else:
            failed += 1
    if failed == 0:
        print("============= PASS ==============")
    else:
        print("============= FAIL ==============")
    print(f"{passed} passed, {failed} failed")


test_cases = submit_cases
if "__RUN__" in globals():
    test_cases = run_cases

main()
