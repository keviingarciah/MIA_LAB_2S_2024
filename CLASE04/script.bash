mkdisk -size=5 -unit=M -fit=WF -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"

fdisk -size=1 -type=P -unit=M -fit=BF -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"
fdisk -size=2 -type=P -unit=M -fit=WF -name="Particion2" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"

mount -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"
mount -name="Particion2" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"

mkfs -id=531A -type=full 
mkfs -id=532A 

rep -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE04/disks/DiscoLab.mia"

