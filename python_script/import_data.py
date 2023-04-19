import psycopg2
import requests
from datetime import datetime, timezone


def get_food_trucks():
    url = 'https://data.sfgov.org/resource/rqzj-sfat.json'
    response = requests.get(url)    
    data = response.json()
    return data

def connect_to_db():
    conn = psycopg2.connect(host="localhost", port="5003", database="elc_db", user="postgres", password="postgres")
    return conn

def process_food_trucks():
    print("here")
    #fetch data
    #for each food_truck
        #transform data
        #insert data to 
    #commit
    #close

    conn = connect_to_db()
    cur = conn.cursor()
    food_trucks = get_food_trucks()

    #get list of all location ids from db
    location_ids = get_location_ids()

    for food_truck in food_trucks:
        dt = datetime.now()

        # dont process&&insert food_trucks that already exists in the db.
        if food_truck['objectid'] in location_ids:
            continue
        status = ""
        facilitytype = ""
        locationdescription = ""
        fooditems = ""
        latitude = ""
        longitude = ""
        if "status" in food_truck:
            #print(f"food_truck [status] : {food_truck['status']}")
            status = food_truck['status']
        if "facilitytype" in food_truck:
            #print(f"food_truck [facilitytype] : {food_truck['facilitytype']}")
            facilitytype = food_truck['facilitytype']
        if "locationdescription" in food_truck:
            # print(f"food_truck [locationdescription] : {food_truck['locationdescription']}")
            locationdescription = food_truck['locationdescription']
        if "fooditems" in food_truck:
            # print(f"food_truck [fooditems] : {food_truck['fooditems']}")
            fooditems = food_truck['fooditems']
        if "latitude" in food_truck:
            # print(f"food_truck [latitude] : {food_truck['latitude']}")
            latitude = food_truck['latitude']
        if "longitude" in food_truck:
            # print(f"food_truck [longitude] : {food_truck['longitude']}")
            longitude = food_truck['longitude']

        # print("\n")
        # print("\n")
        # print("\n")
        insert_query = '''
            INSERT INTO food_trucks (location_id, name, address, status, facility_type, 
            location_description, food_items, latitude, longitude, created_at) 
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            '''
        cur.execute(insert_query, (food_truck['objectid'], food_truck['applicant'], food_truck['address'], status,
                    facilitytype, locationdescription, fooditems, latitude, longitude, dt))
    conn.commit()
    conn.close()

def get_location_ids():
    conn = connect_to_db()
    cur = conn.cursor()
    cur.execute("SELECT location_id FROM food_trucks")
    result = []
    for row in cur.fetchall():
        location_id = row[0]
        result.append(str(location_id))
    conn.close()
    return result

def main():
    process_food_trucks()

main()
#py import_data.py#