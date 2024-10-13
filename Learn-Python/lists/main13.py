def trim_strongholds(strongholds):
    del strongholds[0]
    del strongholds[len(strongholds)-2:len(strongholds)]
    return strongholds

