from main import *

run_cases = [
    (
        {
            "entity": {
                "character": {
                    "name": "Sir Galahad",
                    "quests": {
                        "Dragon_Slayer": {
                            "status": "In Progress",
                        },
                        "Goblin_Hunter": {
                            "status": "Completed",
                        },
                    },
                }
            }
        },
        "In Progress",
    ),
    (
        {
            "entity": {
                "character": {
                    "name": "Lady Gwen",
                    "quests": {
                        "Dragon_Slayer": {
                            "status": "Completed",
                        },
                        "Goblin_Hunter": {
                            "status": "In Progress",
                        },
                    },
                }
            }
        },
        "Completed",
    ),
]

submit_cases = run_cases + [
    (
        {
            "entity": {
                "character": {
                    "name": "Archer Finn",
                    "quests": {
                        "Dragon_Slayer": {
                            "status": "Not Started",
                        },
                        "Goblin_Hunter": {
                            "status": "Completed",
                        },
                    },
                }
            }
        },
        "Not Started",
    ),
    (
        {
            "entity": {
                "character": {
                    "name": "Mage Elara",
                    "quests": {
                        "Dragon_Slayer": {
                            "status": "Failed",
                        },
                        "Goblin_Hunter": {
                            "status": "Completed",
                        },
                    },
                }
            }
        },
        "Failed",
    ),
    (
        {
            "entity": {
                "character": {
                    "name": "Rogue Talon",
                    "quests": {
                        "Dragon_Slayer": {
                            "status": "Completed",
                        },
                        "Goblin_Hunter": {
                            "status": "Not Started",
                        },
                    },
                }
            }
        },
        "Completed",
    ),
]


def test(input1, expected_output):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * Progress Dictionary: {input1}")
    print(f"Expecting: {expected_output}")
    result = get_quest_status(input1)
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
