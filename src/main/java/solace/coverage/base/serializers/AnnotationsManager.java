package solace.coverage.base.serializers;

import java.lang.reflect.Field;

import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;

import solace.coverage.annotations.InnerMetric;
import solace.coverage.base.IMetric;

final class AnnotationsManager {
    public static final JsonObject getInnerMetricsFields(IMetric<?> metric, JsonSerializationContext context) {
        Class<?> clazz = metric.getClass();
        JsonObject inner = new JsonObject();
        boolean exist = false;
        for (Field metricField : clazz.getDeclaredFields()) {
            if (metricField.isAnnotationPresent(InnerMetric.class)) {
                String resultName = metricField.getAnnotation(InnerMetric.class).name();
                if (resultName.isBlank())
                    resultName = metricField.getName();
                //Get field value    
                try {
                    metricField.setAccessible(true);
                    inner.add(resultName, context.serialize(metricField.get(metric)));
                    exist = true;
                } catch (IllegalAccessException e) {
                    System.out.println(e);
                }
            }
        }
        return (exist) ? inner : null;
    }
}