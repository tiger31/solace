package solace.coverage.base.serializers;

import java.lang.reflect.Type;
import java.util.HashMap;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;

import solace.coverage.base.IMapMetric;

public class MapMetricSerializer extends BaseSerializer implements JsonSerializer<IMapMetric<?, ?>> {
    @Override
    public JsonElement serialize(IMapMetric<?, ?> metric, Type metricType, JsonSerializationContext context) {
        JsonObject obj = (JsonObject) super.serialize(metric, metricType, context);
        obj.add("metrics", context.serialize(new HashMap<>(metric)));
        return obj;
    }

}
    