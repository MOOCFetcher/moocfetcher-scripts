import csv
import openpyxl
from openpyxl.styles import Font

wb = openpyxl.Workbook()
ws = wb.active

f = open('courses-annotated.csv')
reader = csv.reader(f, delimiter=',')
for row in reader:
    ws.append(row)
f.close()

wb.save('courses-annotated.xlsx')
