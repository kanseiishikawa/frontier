
import json
from tqdm import tqdm

import sekitoba_library as lib
import sekitoba_data_manage as dm

def json_write( file_name, value ):
    f = open( file_name, "w" )
    json.dump( value, f, ensure_ascii = False, indent = 4 )
    f.close()

def list_max( l ):
    try:
        return max( l )
    except:
        return 0

def main():
    target = {}
    fevalue = {}
    target["Info"] = []
    target["Value"] = []
    target["Info"].append( { "Name": "time",  "Up": False } )
    target["Info"].append( { "Name": "rank",  "Up": False } )
    target["Info"].append( { "Name": "up_time",  "Up": False } )

    fevalue["Info"] = []
    fevalue["Value"] = []
    fevalue["Info"].append( { "Name": "speed" } )
    fevalue["Info"].append( { "Name": "up_speed" } )
    fevalue["Info"].append( { "Name": "pace_speed" } )
    
    race_data = dm.pickle_load( "race_data.pickle" )
    horce_data = dm.dl.data_get( "horce_data_storage.pickle" )
    baba_index_data = dm.dl.data_get( "baba_index_data.pickle" )

    for k in tqdm( race_data.keys() ):
        c += 1
        race_id = lib.id_get( k )
        year = race_id[0:4]
        race_place_num = race_id[4:6]
        day = race_id[9]
        num = race_id[7]

        for kk in race_data[k].keys():
            horce_name = kk.replace( " ", "" )
            current_data, past_data = lib.race_check( horce_data[horce_name],
                                                          year, day, num, race_place_num )#今回と過去のデータに分ける
            cd = lib.current_data( current_data )
            pd = lib.past_data( past_data, current_data )
        
            if not cd.race_check():
                continue

            target_instance = []
            fevalue_instance = []

            speed, up_speed, pace_speed = pd.speed_index( baba_index_data[horce_name] )
            target_instance.append( cd.race_time() )
            target_instance.append( cd.rank() )
            target_instance.append( cd.up_time() )
            
            if len( speed ) == 0:
                continue
            
            fevalue_instance.append( max( speed ) )
            fevalue_instance.append( max( up_speed ) )
            fevalue_instance.append( max( pace_speed ) )

            target["Value"].append( target_instance )
            fevalue["Value"].append( fevalue_instance )

    json_write( "sample_target.json", target )
    json_write( "sample_fevalue.json", fevalue )

main()
