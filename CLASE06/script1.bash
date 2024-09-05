# Crea un disco de 5MB con ajuste Worst Fit
mkdisk -size=5 -unit=M -fit=WF -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE06/disks/DiscoLab.mia"

fdisk -size=1 -type=P -unit=M -fit=BF -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE06/disks/DiscoLab.mia"

mount -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE06/disks/DiscoLab.mia"

mkfs -id=531A 

rep -id=531A -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE06/output/report_mbr.png" -name=mbr
