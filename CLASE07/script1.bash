mkdisk -size=5 -unit=M -fit=WF -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE07/disks/DiscoLab.mia"

fdisk -size=1 -type=P -unit=M -fit=BF -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE07/disks/DiscoLab.mia"

mount -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE07/disks/DiscoLab.mia"

mkfs -id=531A 

mkdir -path="/home"
mkdir -path="/home/usac"
mkdir -path="/home/work"
mkdir -path="/home/usac/mia"

mkfile -size=68 -path=/home/usac/mia/a.txt

rep -id=531A -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE07/output/report_inode.png" -name=inode