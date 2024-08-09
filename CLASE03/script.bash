mkdisk -size=5 -unit=M -fit=WF -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE03/disks/DiscoLab.mia"
fdisk -size=1 -type=L -unit=M -fit=BF -name="Particion3" -path="/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE03/disks/DiscoLab.mia"

# Ejemplo de como ejecutar el script para la Tarea 2
execute -path="script.bash"