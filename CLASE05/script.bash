mkdisk -size=5 -unit=M -fit=WF -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/disks/DiscoLab.mia"

fdisk -size=1 -type=P -unit=M -fit=BF -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/disks/DiscoLab.mia"
fdisk -size=10 -type=P -unit=K -fit=WF -name="Particion2" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/disks/DiscoLab.mia"

mount -name="Particion1" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/disks/DiscoLab.mia"
mount -name="Particion2" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/disks/DiscoLab.mia"

mkfs -id=531A 
mkfs -id=532A -type=full

rep -id=531A -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/output/report_mbr.png" -name=mbr
rep -id=531A -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/output/report_inode.png" -name=inode
rep -id=531A -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE05/output/report_bm_inode.txt" -name=bm_inode
