package item

import "fmt"

type Item uint16

func (i Item) String() string {
	return fmt.Sprintf("0x%04x", uint16(i))
}

const (
	ItemTome1                     Item = 0x0001
	ItemTome2                     Item = 0x0002
	ItemTome3                     Item = 0x0003
	ItemBlessedTome1              Item = 0x005e
	ItemBlessedTome2              Item = 0x005f
	ItemBlessedTome3              Item = 0x0060
	ItemSlimeEssence              Item = 0x0004
	ItemCrystalOozeEssence        Item = 0x0005
	ItemGelatinousChampionEssence Item = 0x0006
	ItemMasteryEssence1           Item = 0x0061
	ItemMasteryEssence2           Item = 0x0062
	ItemMasteryEssence3           Item = 0x0063
	ItemShardPack3                Item = 0x0009
	ItemHonorBadgePack3           Item = 0x000c
	ItemGoldPack3                 Item = 0x003c
	ItemManaPack3                 Item = 0x003f
	ItemPileOfGems                Item = 0x0046
	ItemSentinelSeal              Item = 0x08c6
	ItemWorkHammer3               Item = 0x0011
	ItemWorkHammer4               Item = 0x0012
	ItemWorkHammer5               Item = 0x0013
	ItemWorkHammer6               Item = 0x0014
	ItemEquipmentEvoStone         Item = 0x0016
	ItemEventCoin                 Item = 0x07ec
	ItemBackgroundCoupon          Item = 0x0025
	ItemFortunaToken              Item = 0x002a
	ItemNameEraser                Item = 0x002c
	ItemArchidemonEntryCard       Item = 0x002f
	ItemTalentCard                Item = 0x0033
	ItemQuestCompletionCard       Item = 0x0035
	ItemQuestRefreshCard          Item = 0x0036
	ItemMonsterPass               Item = 0x0038
	ItemTrialEntryCard            Item = 0x0039
	ItemTeamDungeonCard           Item = 0x0043
	ItemBronzeKey                 Item = 0x0047
	ItemSilverKey                 Item = 0x0048
	ItemGoldKey                   Item = 0x0049
	ItemBlueCrystalBoxS           Item = 0x004c
	ItemBlueCrystalBoxL           Item = 0x004d
	ItemRedCrystalBoxL            Item = 0x0051
	ItemStaminaCard               Item = 0x0056
	ItemTeamHBMCard               Item = 0x0057
	ItemLv6TalentRune             Item = 0x0058
	ItemLv7TalentRune             Item = 0x0059
	ItemLv8TalentRune             Item = 0x005a
	ItemLv9TalentRune             Item = 0x08be
	ItemLv10TalentRune            Item = 0x08cd
	ItemTalentChest               Item = 0x005b
	ItemEvolutionRune             Item = 0x005d
	ItemArenaCard                 Item = 0x0031
)
